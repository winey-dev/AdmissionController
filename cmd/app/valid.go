package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"k8s.io/api/admission/v1beta1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	ValidationAnnotations = "yiaw.webhook/validation"
)

func Validating(ar *v1beta1.AdmissionReview) *v1beta1.AdmissionResponse {
	req := ar.Request
	var pod corev1.Pod
	if err := json.Unmarshal(req.Object.Raw, &pod); err != nil {
		log.Printf("Could not unmarshal raw object: %v", err)
		return &v1beta1.AdmissionResponse{
			Result: &metav1.Status{
				Message: err.Error(),
			},
		}
	}

	annotations := pod.ObjectMeta.GetAnnotations()
	if annotations == nil {
		return &v1beta1.AdmissionResponse{
			Allowed: true,
		}
	}

	var resp *v1beta1.AdmissionResponse
	switch strings.ToLower(annotations[ValidationAnnotations]) {
	case "yes", "true", "ok":
		resp = &v1beta1.AdmissionResponse{
			Allowed: false,
			Result: &metav1.Status{
				Message: "Valid Check, Not Allowed Pod Create",
			},
		}
	default:
		resp = &v1beta1.AdmissionResponse{
			Allowed: true,
		}
	}

	return resp
}

func ValidatingWebHook(w http.ResponseWriter, r *http.Request) {
	var body []byte
	if r.Body != nil {
		if data, err := ioutil.ReadAll(r.Body); err == nil {
			body = data
		}
	}
	if len(body) == 0 {
		log.Println("empty body")
		http.Error(w, "empty body", http.StatusBadRequest)
		return
	}

	log.Printf("Recv Message : %s\n", string(body))

	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" {
		log.Printf("Content-Type=%s, expect application/json", contentType)
		http.Error(w, "invalid Content-Type, expect `application/json`", http.StatusUnsupportedMediaType)
		return
	}

	var admissionResponse *v1beta1.AdmissionResponse
	ar := v1beta1.AdmissionReview{}
	if _, _, err := deserializer.Decode(body, nil, &ar); err != nil {
		admissionResponse = &v1beta1.AdmissionResponse{
			Result: &metav1.Status{
				Message: err.Error(),
			},
		}
	} else {
		admissionResponse = Validating(&ar)
	}

	responseReview := v1beta1.AdmissionReview{}
	responseReview.Response = admissionResponse

	resp, err := json.Marshal(responseReview)
	if err != nil {
		http.Error(w, fmt.Sprintf("json Marshal Fail.. err=%v", err), http.StatusInternalServerError)
		return
	}

	log.Printf("Send Message : %s\n", string(resp))
	w.Write(resp)
}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	v1 "k8s.io/api/admission/v1"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		ar := v1.AdmissionReview{}
		if err := json.NewDecoder(r.Body).Decode(&ar); err != nil {
			panic(err)
		}

		for key, values := range r.Header {
			fmt.Printf("%s: %v\n", key, values)
		}
		fmt.Printf("%+v", ar)
		fmt.Println()

		admissionReview := v1.AdmissionReview{
			TypeMeta: ar.DeepCopy().TypeMeta,
			Response: &v1.AdmissionResponse{
				Allowed: true,
				UID:     ar.Request.UID,
			},
		}

		data, err := json.Marshal(admissionReview)
		if err != nil {
			panic(err)
		}

		if n, err := w.Write(data); err != nil {
			panic(err)
		} else {
			fmt.Println(n)
		}
	})

	keyfile := os.Getenv("KEYFILE")
	cafile := os.Getenv("CERTFILE")
	err := http.ListenAndServeTLS(":443", cafile, keyfile, nil)
	if err != nil {
		fmt.Println(err)
	}
}

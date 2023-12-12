# Technologies

* Fiber framework.
* DialogFlow NLP.
* Kommunicate.io.
* Redis
* Docker
* GCP

# Goal

Offer a simple demo promotional chatbot flow that offers coupons to prospecting customers. The chosen programming language for the solution is Golang.

# Architecture
```
.
├── build
├── cmd
└── internal
    ├── config
    ├── controller
    │   ├── coupon
    │   └── webhook
    ├── model
    │   ├── controller
    │   └── service
    ├── pkg
    │   ├── chatbot
    │   ├── redis
    │   └── util
    ├── routes
    │   ├── coupon
    │   └── webhook
    └── service
        ├── chatbot
        ├── coupon
        └── webhook
```

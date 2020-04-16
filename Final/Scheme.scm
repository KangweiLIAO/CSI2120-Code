#lang scheme

(define names
     '("marie" "jean" "claude" "emma" "sam" "tom" "eve" "bob"))

(define (first n lst)
    (if (or (null? lst) (= n 0))    
        '()    ; if list = null return '()
        (cons (car lst) (first (- n 1) (cdr lst)))  ; loop the list until n = 0/empty
    )
)

(define (insertAt e lst i)
    (cond 
        ((null? lst) '())               ; if list = null return '()
        ((zero? i) (cons e lst))        ; i = 0 indicates we should insert element here
        (else (cons (car lst)(insertAt e (cdr lst) (- i 1))))   ; else loop the list until i = 0/empty 
    )
)

(define (shuffle lst n)
    (if (zero? n) 
        lst
        (begin 
            (insertAt (car names) (cdr names) (random (- (length names) 1)))
            (shuffle lst (- n 1))
        ) 
    )
)

(define (winner lst n)
    (first n (shuffle lst 20))
)
kubectl run -i --tty loadgen --rm --image=radial/busyboxplus:curl --restart=Never -- /bin/sh -c "while true; do curl -sS http://healthapp-service:80/ ; sleep 0.01 ; done"

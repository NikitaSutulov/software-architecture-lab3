startPos=25
finishPos=100
x=$startPos
y=$startPos
step=10
interval=0.01

curl -X POST http://localhost:17000 -d "white"

curl -X POST http://localhost:17000 -d "figure $(awk -v s=$startPos 'BEGIN{printf "%.2f %.2f", s/100, s/100}')"

curl -X POST http://localhost:17000 -d "update"

sleep $interval

while true; do
  while ((x < finishPos-startPos)); do
    curl -X POST http://localhost:17000 -d "move $(awk -v s=$step 'BEGIN{printf "%.2f 0", s/100}')"
    x=$((x + step))
    curl -X POST http://localhost:17000 -d "update"
    sleep $interval
  done

  while ((y < finishPos-startPos)); do
    curl -X POST http://localhost:17000 -d "move $(awk -v s=$step 'BEGIN{printf "0 %.2f", s/100}')"
    y=$((y + step))
    curl -X POST http://localhost:17000 -d "update"
    sleep $interval
  done

  while ((x > startPos)); do
    curl -X POST http://localhost:17000 -d "move $(awk -v s=$step 'BEGIN{printf "%.2f 0", -s/100}')"
    x=$((x - step))
    curl -X POST http://localhost:17000 -d "update"
    sleep $interval
  done

  while ((y > startPos)); do
    curl -X POST http://localhost:17000 -d "move $(awk -v s=$step 'BEGIN{printf "0 %.2f", -s/100}')"
    y=$((y - step))
    curl -X POST http://localhost:17000 -d "update"
    sleep $interval
  done
done

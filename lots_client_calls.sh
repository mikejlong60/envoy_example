counter=1
while [ $counter -le 1000 ]
do
echo $counter
  ((counter++))
  go run ./client/main.go
done
echo All done

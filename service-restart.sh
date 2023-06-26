#restart mysql
sudo systemctl restart mysql

sudo service server restart

while IFS='=' read -r key value; do
  if [ "$key" = "MO_THREAD" ]; then
    n=$value

    i=1
    while [ $i -le $n ]
    do
      sudo service consumer-mo@"thread_$i" restart
      sleep 1
      i=$((i+1))
    done
  fi
done < "app.env"

while IFS='=' read -r key value; do
  if [ "$key" = "DR_THREAD" ]; then
    n=$value

    i=1
    while [ $i -le $n ]
    do
      sudo service consumer-dr@"thread_$i" restart
      sleep 1
      i=$((i+1))
    done
  fi
done < "app.env"

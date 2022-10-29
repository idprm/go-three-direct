sudo service server stop

sudo service publisher-renewal stop
sudo service publisher-retry stop

while IFS='=' read -r key value; do
  if [ "$key" = "RENEWAL_THREAD" ]; then
    n=$value

    i=1
    while [ $i -le $n ]
    do
      sudo service consumer-renewal@"thread_$i" stop
      sleep 1
      i=$((i+1))
    done
  fi
done < "app.env"


while IFS='=' read -r key value; do
  if [ "$key" = "RETRY_THREAD" ]; then
    n=$value

    i=1
    while [ $i -le $n ]
    do
      sudo service consumer-retry@"thread_$i" stop
      sleep 1
      i=$((i+1))
    done
  fi
done < "app.env"

while IFS='=' read -r key value; do
  if [ "$key" = "PURGE_THREAD" ]; then
    n=$value

    i=1
    while [ $i -le $n ]
    do
      sudo service consumer-purge@"thread_$i" stop
      sleep 1
      i=$((i+1))
    done
  fi
done < "app.env"

while IFS='=' read -r key value; do
  if [ "$key" = "MO_THREAD" ]; then
    n=$value

    i=1
    while [ $i -le $n ]
    do
      sudo service consumer-mo@"thread_$i" stop
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
      sudo service consumer-dr@"thread_$i" stop
      sleep 1
      i=$((i+1))
    done
  fi
done < "app.env"
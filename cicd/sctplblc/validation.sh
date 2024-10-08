#!/bin/bash
source ../common.sh
echo SCENARIO-sctplb-lc

servArr=( "server1" "server2" "server3" )
ep=( "31.31.31.1" "32.32.32.1" "33.33.33.1" )

$hexec l3ep1 ../common/sctp_server ${ep[0]} 8080 server1 >/dev/null 2>&1 &
$hexec l3ep2 ../common/sctp_server ${ep[1]} 8080 server2 >/dev/null 2>&1 &
$hexec l3ep3 ../common/sctp_server ${ep[2]} 8080 server3 >/dev/null 2>&1 &

sleep 5
code=0
j=0
waitCount=0
while [ $j -le 2 ]
do
    res=$($hexec l3h1 timeout 10 ../common/sctp_client 10.10.10.1 0 ${ep[j]} 8080)
    #echo $res
    if [[ $res == "${servArr[j]}" ]]
    then
        echo "$res UP"
        j=$(( $j + 1 ))
    else
        echo "Waiting for ${servArr[j]}(${ep[j]})"
        waitCount=$(( $waitCount + 1 ))
        if [[ $waitCount == 10 ]];
        then
            echo "All Servers are not UP"
            echo SCENARIO-sctplb-lc [FAILED]
            sudo pkill sctp_server >/dev/null 2>&1
            exit 1
        fi

    fi
    sleep 1
done

for i in {1..2}
do
for j in {0..2}
do
    res=$($hexec l3h1 timeout 10 ../common/sctp_client 10.10.10.1 0 20.20.20.1 2020)
    echo -e $res
    if [[ $res != "${servArr[0]}" ]]
    then
        code=1
    fi
    sleep 1
done
done

echo "Holding connection at"
$hexec l3h1 ../common/sctp_client 10.10.10.1 0 20.20.20.1 2020 20 &
sleep 1

echo -e "\n\nTesting Service IP: 20.20.20.1"
code=0
for i in {1..2}
do
for j in {0..2}
do
    res=$($hexec l3h1 timeout 10 ../common/sctp_client 10.10.10.1 0 20.20.20.1 2020)
    echo -e $res
    if [[ $res != "${servArr[1]}" ]]
    then
        code=1
    fi
    sleep 1
done
done

sudo pkill sctp_server >/dev/null 2>&1
sudo killall -9 ncat >/dev/null 2>&1
rm -f nohup.out >/dev/null 2>&1
if [[ $code == 0 ]]
then
    echo SCENARIO-sctplb-lc [OK]
else
    echo SCENARIO-sctplb-lc [FAILED]
fi
exit $code


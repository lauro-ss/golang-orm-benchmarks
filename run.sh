for i in {1..10};
do
    echo "-----" $i "-----";
    make benchmark-insert; 
done
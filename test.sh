for i in noblanks.txt blanklines.txt doublespaced.txt; do
  if ! diff <(./testmars < data/$i) data/expected.txt > /dev/null; then
    echo Mismatched output for $i
  fi
done

for i in empty.txt; do
  if diff <(./testmars < data/$i) data/expected.txt > /dev/null; then
    echo Mismatched output for $i
  fi
done

echo "Release Notes"
for file in *.yaml
do
    yq read $file ReleaseNotes
done


echo "Upgrade Notes"
for file in *.yaml
do
    yq read $file UpgradeNotes
done

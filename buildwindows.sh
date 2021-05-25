rm -rf publish/
rm mardinlee.zip

mkdir publish
mkdir publish/views
mkdir publish/assets

cp web.config publish/
cp -r views/. publish/views/
cp -r assets/. publish/assets/

GOOS=windows GOARCH=386 go build -o publish/gamebox.exe .

cd ./publish
zip ../mardinlee.zip . -r
cd ..
rm -rf publish
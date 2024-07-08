run: 
	./runner.sh

css:
	npx tailwindcss build -i views/css/main.css -o public/style.css --watch

templ:
	templ generate --watch --proxy="http://localhost:4000" --open-browser=false

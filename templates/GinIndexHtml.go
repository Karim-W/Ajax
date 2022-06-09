package templates

var GinIndexHtml = `<!doctype html>
<html class="w-screen h-screen">

<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<script src="https://cdn.tailwindcss.com"></script>
</head>


<body
	class="w-full h-full bg-gradient-to-br pt-[20rem] from-cyan-500 via-blue-500 via-pink-500 to-red-500 flex flex-col items-center justify-center gap-12">
	<script>
		n = new Date();
		y = n.getFullYear();
		m = n.getMonth() + 1;
		d = n.getDate();
		document.getElementById("date").innerHTML = m + "/" + d + "/" + y;
	</script>
	<img src="https://raw.githubusercontent.com/gin-gonic/logo/master/color.png" class="w-[10rem]" />
	<h1 style="text-shadow: 1px 1px 2px ;"
		class="text-6xl font-bold text-white accent-teal-500 text-transparent bg-clip-text bg-gradient-to-r from-zinc-900 to-zinc-800 underline underline-offset-4">
		Yay! you're on Go!</h1>
	<p class="text-3xl text-zinc-200">
		To view Swagger UI, click <a class="text-cyan-900 underline underline-offset-2"
			href="/swagger/index.html">here</a>
	</p>
	<p class="text-xl text-zinc-200" id="date">
		Viewed on:
		<script> document.write(new Date()); </script>
	</p>

	<p class="text-sm text-zinc-900 mt-auto mb-[2rem]" id="date">
		©
		<script> document.write(new Date().getFullYear()); </script> Karim Hassan. All rights reserved.
	</p>

</body>

</html>`

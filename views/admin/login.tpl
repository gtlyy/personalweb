<!DOCTYPE html>
<meta charset="UTF-8">
<title>登录</title>
<script src="/static/js/tailwindcss.js"></script>
<body class="bg-slate-100 min-h-screen flex items-center justify-center">
<div class="w-full max-w-md bg-white p-8 rounded-xl shadow">
  <h1 class="text-2xl font-bold text-center mb-6">后台登录</h1>
  {{if .Msg}}<div class="text-red-600 text-center mb-4">{{.Msg}}</div>{{end}}
  <form method="post" action="/admin/login">
    <div class="mb-4">
      <label class="block text-gray-700 mb-1">账号</label>
      <input name="username" class="w-full border px-3 py-2 rounded" required>
    </div>
    <div class="mb-6">
      <label class="block text-gray-700 mb-1">密码</label>
      <input type="password" name="password" class="w-full border px-3 py-2 rounded" required>
    </div>
    <button class="w-full bg-indigo-600 text-white py-2 rounded hover:bg-indigo-700">登录</button>
  </form>
</div>
</body>
</html>

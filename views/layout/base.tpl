<!DOCTYPE html>
<html lang="zh-CN">
<head>
<meta charset="UTF-8">
<title>个人主页</title>
<script src="/static/js/tailwindcss.js"></script>
</head>
<body class="bg-slate-50 min-h-screen">
<nav class="bg-white shadow-md">
  <div class="max-w-6xl mx-auto px-4 py-3 flex justify-between items-center">
    <a href="/" class="text-xl font-bold text-indigo-600">My Site</a>
    <div class="space-x-4 text-sm md:text-base">
      <a href="/article" class="text-gray-700 hover:text-indigo-600">文章</a>
      <a href="/games" class="text-gray-700 hover:text-indigo-600">游戏</a>
      <a href="/tools" class="text-gray-700 hover:text-indigo-600">工具</a>
      {{if .Admin}}
      <a href="/admin/index" class="text-indigo-600">管理</a>
      <a href="/admin/game" class="text-indigo-600">游戏管理</a>
      <a href="/admin/tool" class="text-indigo-600">工具管理</a>
      <a href="/admin/password" class="text-indigo-600">修改密码</a>
      <a href="/admin/logout" class="text-red-600">退出</a>
      {{else}}
      <a href="/admin/login" class="text-indigo-600">登录</a>
      {{end}}
    </div>
  </div>
</nav>
<div class="max-w-6xl mx-auto px-4 py-8">
{{template "content" .}}
</div>
</body>
</html>

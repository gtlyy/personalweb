{{template "layout/base.tpl" .}}
{{define "content"}}
<h1 class="text-2xl font-bold mb-6">新增文章</h1>
<div class="bg-white p-6 rounded shadow">
<form method="post">
  <input type="hidden" name="_xsrf" value="{{.xsrf_token}}">
  <div class="mb-4">
    <label class="block mb-1 font-medium">标题</label>
    <input name="title" class="w-full border px-3 py-2 rounded" required>
  </div>
  <div class="mb-4">
    <label class="block mb-1 font-medium">分类</label>
    <input name="category" class="w-full border px-3 py-2 rounded" required>
  </div>
  <div class="mb-4">
    <label class="block mb-1 font-medium">状态</label>
    <label class="mr-4"><input type="radio" name="status" value="1" checked> 草稿</label>
    <label><input type="radio" name="status" value="2"> 发布</label>
  </div>
  <div class="mb-6">
    <label class="block mb-1 font-medium">Markdown 内容</label>
    <link rel="stylesheet" href="/static/css/easymde.min.css">
    <textarea id="content" name="content"></textarea>
  </div>
  <button class="bg-indigo-600 text-white px-4 py-2 rounded">保存</button>
</form>
</div>
<script src="/static/js/easymde.min.js"></script>
<script>
var easyMDE = new EasyMDE({ element: document.getElementById("content") });
</script>
{{end}}

{{template "layout/base.tpl" .}}
{{define "content"}}
<h1 class="text-2xl font-bold mb-6">上传游戏</h1>
<div class="bg-white p-6 rounded shadow">
<form method="post" enctype="multipart/form-data">
  <div class="mb-4">
    <label class="block mb-1 font-medium">游戏名称</label>
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
    <label class="block mb-1 font-medium">上传 ZIP（内含 index.html）</label>
    <input type="file" name="zipfile" accept=".zip" class="w-full border px-3 py-2 rounded" required>
  </div>
  <button class="bg-emerald-600 text-white px-4 py-2 rounded">上传并发布</button>
</form>
</div>
{{end}}

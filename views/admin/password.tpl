{{template "layout/base.tpl" .}}
{{define "content"}}
<h1 class="text-2xl font-bold mb-6">修改密码</h1>
<div class="bg-white p-6 rounded shadow max-w-xl">
  {{if .Msg}}<div class="mb-4 text-red-600">{{.Msg}}</div>{{end}}
<form method="post">
  <input type="hidden" name="_xsrf" value="{{.xsrf_token}}">
  <div class="mb-4">
    <label class="block mb-1 font-medium">当前密码</label>
    <input type="password" name="old_password" class="w-full border px-3 py-2 rounded" required>
  </div>
  <div class="mb-4">
    <label class="block mb-1 font-medium">新密码</label>
    <input type="password" name="new_password" class="w-full border px-3 py-2 rounded" required>
  </div>
  <div class="mb-6">
    <label class="block mb-1 font-medium">确认新密码</label>
    <input type="password" name="confirm_password" class="w-full border px-3 py-2 rounded" required>
  </div>
  <button class="bg-indigo-600 text-white px-5 py-2 rounded">保存修改</button>
</form>
</div>
{{end}}

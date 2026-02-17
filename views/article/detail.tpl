{{template "layout/base.tpl" .}}
{{define "content"}}
<link rel="stylesheet" href="/static/css/article.css">
<div class="max-w-3xl mx-auto bg-white rounded-2xl shadow-sm p-6 md:p-10">
  <span class="inline-block px-2 py-1 text-xs bg-indigo-100 text-indigo-600 rounded-full mb-4">{{.Article.Category}}</span>
  <h1 class="text-3xl md:text-4xl font-bold text-gray-800 mb-4">{{.Article.Title}}</h1>
  <p class="text-gray-400 text-sm mb-8">{{.Article.CreateTime.Format "2006-01-02"}}</p>
  <div class="article-content" id="articleContent">
    {{str2html .Article.ContentMd}}
  </div>
</div>
<script>
(function() {
    var content = document.getElementById('articleContent');
    var elements = content.querySelectorAll('p, h1, h2, h3, h4, h5, h6, ul, ol, blockquote');
    elements.forEach(function(parent) {
        var codes = parent.querySelectorAll('code');
        if (codes.length > 0) {
            codes.forEach(function(code) {
                var text = code.textContent || code.innerText;
                if (text.indexOf('\n') > -1 || text.length > 40) {
                    var pre = document.createElement('pre');
                    var codeInner = document.createElement('code');
                    codeInner.textContent = text.replace(/^javascript\n?/, '');
                    pre.appendChild(codeInner);
                    code.parentNode.replaceChild(pre, code);
                }
            });
        }
    });
})();
</script>
{{end}}

- commitを取り消したいとき

```git
前回のコミットの状態に戻したい時(ファイルの内容も戻ってしまう)
git reset --hard HEAD^
前回のコミットだけ取り消したい時
git reset --soft HEAD^
前回のコミットとaddも取り消したい時
git reset --mixed HEAD^

git reset のわかりやすい説明
https://qiita.com/shuntaro_tamura/items/db1aef9cf9d78db50ffe
```


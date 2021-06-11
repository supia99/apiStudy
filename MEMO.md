# 開発メモ
* [ローカルからブランチを切って、リモートにpushする方法](https://qiita.com/shungo_m/items/4bbcbcbd3d7aea40b21e)
  ~~~
  # 現在のブランチからブランチを切る
  newBranch=new-branch
  git checkout -b $newBranch
  git push --set-upstream origin $newBranch
  ~~~
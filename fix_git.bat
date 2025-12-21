@echo off
git rm -r --cached ad-frontend/.next > rm_next.log 2>&1
git commit -m "chore: stop tracking .next folder" > commit_next.log 2>&1
git push origin main > push_next.log 2>&1

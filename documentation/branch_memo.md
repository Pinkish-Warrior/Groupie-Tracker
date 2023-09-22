## add branch from Local 
  - git branch (branchName)  >>>>then it will create branch under .git/.........
  - git push -u origin branchName   >>push the branch to gittea
  
## swich   
  - git switch (branchName)  >>>>>this commend will swtich to the branch or master
  - git swtich master  >>>swich to master 

## push changed file from branch (need to be make sure still on same branch)
  - git add filename 
  - git commit -m "like normal"
  - git push 
  
## To merge (*need to the place(master/branchName))  ex:merge branch to master, then switch to master first
  - git switch master 
  - git merge branchName   >>merge master with branch 
  
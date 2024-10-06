

FOR /L %%A IN (1,1,10) DO (
  ECHO %%A
  dir >> %%A".txt"
)
FOR /L %%A IN (1,1,10) DO (
  MD %%A
  FOR /L %%B IN (1,1,20) DO (
    dir >> %%A/%%B.txt
  )
)
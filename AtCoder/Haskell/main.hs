import           Control.Monad
import           Data.List
import           Data.String

main = do
  n <- fmap read getLine
  [x, y, z] <- fmap readWords getLine
  lines <- replicateM n getLine
  let [a, b] = (transpose . map readWords) lines
  print $ solve a b x y z
  where
    readWords = map read . words

solve :: [Int] -> [Int] -> Int -> Int -> Int -> Int
solve a b x y z = sum a + sum b + x + y + z

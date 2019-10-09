import           Control.Monad
import qualified Data.ByteString.Char8 as BS
import           Data.Char
import           Data.List
import           Data.Maybe

main = do
  n <- fmap (fst . fromJust . BS.readInt) BS.getLine
  [x, y, z] <- readInts <$> BS.getLine
  lines <- replicateM n BS.getLine
  let [a, b] = transpose $ map readInts lines
  print $ solve a b x y z
  where
    readInts :: BS.ByteString -> [Int]
    readInts = unfoldr (BS.readInt . BS.dropWhile isSpace)

solve :: [Int] -> [Int] -> Int -> Int -> Int -> Int
solve a b x y z = sum a + sum b + x + y + z

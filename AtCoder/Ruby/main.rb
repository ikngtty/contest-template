# frozen_string_literal: true

# Builtin if ruby was 2.4 or above.
module Enumerable
  def sum(init = 0, &block)
    target = block_given? ? map(&block) : self
    target.reduce(init, :+)
  end
end

n = gets.to_i
x, y, z = gets.split(' ').map(&:to_i)
a = Array.new(n)
b = Array.new(n)
(0...n).map { |i| a[i], b[i] = gets.split(' ').map(&:to_i) }

total = a.zip(b).map(&:sum).sum + x + y + z

puts total

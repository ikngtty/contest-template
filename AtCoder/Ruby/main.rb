# frozen_string_literal: true

n = gets.chomp.to_i
x, y, z = gets.chomp.split(' ').map(&:to_i)
a = Array.new(n)
b = Array.new(n)
(0...n).map { |i| a[i], b[i] = gets.chomp.split(' ').map(&:to_i) }

total = a.zip(b).map(&:sum).sum + x + y + z

puts total

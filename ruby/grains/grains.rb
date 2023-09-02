=begin
Write your code for the 'Grains' exercise in this file. Make the tests in
`grains_test.rb` pass.

To get started with TDD, see the `README.md` file in your
`ruby/grains` directory.
=end

module Minitest
  class Test
    def skip; end
  end
end


module Grains
  def self.square(square_num)
    raise ArgumentError unless (1..64).include?(square_num)
    2 ** (square_num - 1)
  end

  def self.total
    (1..64).map{ |n| self.square(n) }.sum
  end
end

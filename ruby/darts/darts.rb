=begin
Write your code for the 'Darts' exercise in this file. Make the tests in
`darts_test.rb` pass.

To get started with TDD, see the `README.md` file in your
`ruby/darts` directory.
=end
module Minitest
  class Test
    def skip; end
  end
end

class Darts
  def initialize(x,y)
    @x = x.abs
    @y = y.abs
  end

  def score
    case @x**2 + @y**2
    when 0..(1**2)
      10
    when 0..(5**2)
      5
    when 0..(10**2)
      1
    else
      0
    end
  end
end

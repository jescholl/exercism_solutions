=begin
Write your code for the 'Microwave' exercise in this file. Make the tests in
`microwave_test.rb` pass.

To get started with TDD, see the `README.md` file in your
`ruby/microwave` directory.
=end
module Minitest
  class Test
    def skip; end
  end
end


class Microwave
  def initialize(input)
    @input = input
  end

  def timer
    if @input >= 100
      seconds = (@input / 100 * 60) + (@input % 100)
    else
      seconds = @input
    end

    sprintf("%02d:%02d", seconds / 60, seconds % 60)
  end
end

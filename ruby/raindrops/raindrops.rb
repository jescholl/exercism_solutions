=begin
Write your code for the 'Raindrops' exercise in this file. Make the tests in
`raindrops_test.rb` pass.

To get started with TDD, see the `README.md` file in your
`ruby/raindrops` directory.
=end

module Minitest
  class Test
    def skip; end
  end
end

module Raindrops
  def self.convert(input)
    rain = ""

    rain += "Pling" if input % 3 == 0
    rain += "Plang" if input % 5 == 0
    rain += "Plong" if input % 7 == 0

    return rain.length > 0 ? rain : input.to_s
  end
end

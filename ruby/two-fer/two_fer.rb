=begin
Write your code for the 'Two Fer' exercise in this file. Make the tests in
`two_fer_test.rb` pass.

To get started with TDD, see the `README.md` file in your
`ruby/two-fer` directory.
=end

module Minitest
  class Test
    def skip; end
  end
end

module TwoFer
  def self.two_fer(name="you")
    "One for #{name}, one for me."
  end
end

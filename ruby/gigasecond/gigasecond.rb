=begin
Write your code for the 'Gigasecond' exercise in this file. Make the tests in
`gigasecond_test.rb` pass.

To get started with TDD, see the `README.md` file in your
`ruby/gigasecond` directory.
=end

module Minitest
  class Test
    def skip; end
  end
end

module Gigasecond
  def self.from(time)
    time + 10**9
  end
end

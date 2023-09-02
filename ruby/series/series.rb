=begin
Write your code for the 'Series' exercise in this file. Make the tests in
`series_test.rb` pass.

To get started with TDD, see the `README.md` file in your
`ruby/series` directory.
=end

module Minitest
  class Test
    def skip; end
  end
end

class Series
  def initialize(str)
    @str = str
  end

  def slices(slices)
    if slices > @str.length
      raise ArgumentError
    end
    @str.length.-(slices - 1).times.map { |i| @str[i,slices] }
  end
end

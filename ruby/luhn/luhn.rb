=begin
Write your code for the 'Luhn' exercise in this file. Make the tests in
`luhn_test.rb` pass.

To get started with TDD, see the `README.md` file in your
`ruby/luhn` directory.
=end

module Minitest
  class Test
    def skip; end
  end
end


module Luhn
  def self.valid?(str)
    str.gsub!(' ', '')
    if str.length <= 1 || str.match?(/\D/)
      return false
    end

    arr = str.chars.map(&:to_i)

    arr.reverse!.map!.with_index do |n,i|
      if i%2 == 1
        n *= 2
        n -= 9 if n > 9
      end
      n
    end

    return arr.inject(&:+) % 10 == 0
  end
end

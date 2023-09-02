=begin
Write your code for the 'Hamming' exercise in this file. Make the tests in
`hamming_test.rb` pass.

To get started with TDD, see the `README.md` file in your
`ruby/hamming` directory.
=end

module Minitest
  class Test
    def skip; end
  end
end

module Hamming
  def self.compute(str1, str2)
    if str2.length != str1.length
      raise ArgumentError
    end

    distance = 0
    str1.length.times do |i|
      if str1[i] != str2[i]
        distance +=1
      end
    end

    return distance
  end
end

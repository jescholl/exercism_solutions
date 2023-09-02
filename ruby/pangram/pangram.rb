=begin
Write your code for the 'Pangram' exercise in this file. Make the tests in
`pangram_test.rb` pass.

To get started with TDD, see the `README.md` file in your
`ruby/pangram` directory.
=end

module Minitest
  class Test
    def skip; end
  end
end


module Pangram
  @letters = %w{a b c d e f g h i j k l m n o p q r s t u v w x y z}
  def self.pangram?(sentence)
    @letters.each do |letter|
      return false unless sentence.downcase.include?(letter)
    end
    return true
  end
end

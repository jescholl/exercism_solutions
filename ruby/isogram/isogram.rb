=begin
Write your code for the 'Isogram' exercise in this file. Make the tests in
`isogram_test.rb` pass.

To get started with TDD, see the `README.md` file in your
`ruby/isogram` directory.
=end

module Minitest
  class Test
    def skip; end
  end
end


module Isogram
  def self.isogram?(input)
    input.downcase.gsub(/[ -]/, '').chars.group_by{ |c| c }.all?{ |_char,occurences| occurences.length == 1 }
  end
end

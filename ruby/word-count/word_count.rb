=begin
Write your code for the 'Word Count' exercise in this file. Make the tests in
`word_count_test.rb` pass.

To get started with TDD, see the `README.md` file in your
`ruby/word-count` directory.
=end

module Minitest
  class Test
    def skip; end
  end
end


class Phrase
  def initialize(words)
    @words = words
  end

  def word_count
    result = {}
    @words.downcase.gsub(/ ['"]|['"] /, ' ').gsub(/[a-zA-Z0-9']+/).each do |word|
      result[word] ||= 0
      result[word] += 1
    end

    return result
  end

end

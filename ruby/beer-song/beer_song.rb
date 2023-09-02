=begin
Write your code for the 'Beer Song' exercise in this file. Make the tests in
`beer_song_test.rb` pass.

To get started with TDD, see the `README.md` file in your
`ruby/beer-song` directory.
=end

module Minitest
  class Test
    def skip; end
  end
end

module BeerSong

  def self.verse(num)
    song = "#{ bottles(num).capitalize } on the wall, #{ bottles(num) }.\n"

    if num == 0
      song += "Go to the store and buy some more"
      num = 100
    else
      song += "Take #{ num == 1 ? 'it' : 'one' } down and pass it around"
    end

    song + ", #{ bottles(num-1) } on the wall.\n"
  end

  def self.bottles(num)
    "#{num == 0 ? 'no more' : num } bottle#{ num == 1 ? '' : 's' } of beer"
  end

  def self.recite(first_verse, verse_count)
    first_verse.downto(first_verse - verse_count + 1).map.with_index do |verse_num, i|
      verse(verse_num)
    end.join("\n")
  end
end

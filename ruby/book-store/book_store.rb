=begin
Write your code for the 'Book Store' exercise in this file. Make the tests in
`book_store_test.rb` pass.

To get started with TDD, see the `README.md` file in your
`ruby/book-store` directory.
=end

module Minitest
  class Test
    def skip(_msg='', _bt=caller)
    end
  end
end

module BookStore
  @base_price = 8

  def self.calculate_price(basket)
    return [
      price_with_max_group_size(basket),
      price_with_max_group_count(basket)
    ].min
  end

  def self.price_with_max_group_size(basket)
    max_count = basket.map { |book| basket.count(book) }.max || 0
    sets = max_count.times.map { [] }

    basket.each do |book|
      sets.each do |set|
        if !set.include?(book)
          set << book
          break
        end
      end
    end

    return sets.map { |set| set_price(set.compact.length) }.inject(&:+) || 0
  end

  def self.price_with_max_group_count(basket)
    max_count = basket.map { |book| basket.count(book) }.max || 0
    sets = max_count.times.map { [] }

    basket.each do |book|
      sets.each do |set|
        if !set.include?(book) and set.length == sets.map(&:length).min
          set << book
          break
        end
      end
    end

    return sets.map { |set| set_price(set.compact.length) }.inject(&:+) || 0
  end

  def self.set_price(set_size)
    [
      1,
      1,
      0.95,
      0.9,
      0.8,
      0.75,
    ][set_size] * set_size * @base_price
  end

end

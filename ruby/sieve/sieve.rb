=begin
Write your code for the 'Sieve' exercise in this file. Make the tests in
`sieve_test.rb` pass.

To get started with TDD, see the `README.md` file in your
`ruby/sieve` directory.
=end

module Minitest
  class Test
    def skip; end
  end
end


class Sieve
  def initialize(max)
    @max = max
  end

  def calc_primes
    @primes = []
    candidates = Array.new(@max + 1, true)
    (2..@max).each do |num|
      next unless candidates[num]

      @primes << num

      (num..@max).step(num) do |multiple|
        candidates[multiple] = false
      end
    end
  end

  def primes
    @primes || calc_primes && @primes
  end
end

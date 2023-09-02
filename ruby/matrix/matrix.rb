=begin
Write your code for the 'Matrix' exercise in this file. Make the tests in
`matrix_test.rb` pass.

To get started with TDD, see the `README.md` file in your
`ruby/matrix` directory.
=end


class Matrix
  def initialize(input)
    @matrix = input.split("\n").map do |line|
      line.split(" ").map(&:to_i)
    end
  end

  def rows
    @matrix
  end

  def columns
    @matrix.transpose
  end
end

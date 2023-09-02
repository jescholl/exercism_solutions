=begin
Write your code for the 'Rna Transcription' exercise in this file. Make the tests in
`rna_transcription_test.rb` pass.

To get started with TDD, see the `README.md` file in your
`ruby/rna-transcription` directory.
=end

module Minitest
  class Test
    def skip; end
  end
end

module Complement

  @dna_rna_map = {
    'G' => 'C',
    'C' => 'G',
    'T' => 'A',
    'A' => 'U',
  }

  def self.of_dna(dna)
    dna.chars.map{ |nucleotide| @dna_rna_map[nucleotide] }.join
  end
end

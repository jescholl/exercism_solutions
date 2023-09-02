=begin
Write your code for the 'Twelve Days' exercise in this file. Make the tests in
`twelve_days_test.rb` pass.

To get started with TDD, see the `README.md` file in your
`ruby/twelve-days` directory.
=end


module TwelveDays
  @days = [
    { english: "first", gift: "a Partridge in a Pear Tree" },
    { english: "second", gift: "two Turtle Doves" },
    { english: "third", gift: "three French Hens" },
    { english: "fourth", gift: "four Calling Birds" },
    { english: "fifth", gift: "five Gold Rings" },
    { english: "sixth", gift: "six Geese-a-Laying" },
    { english: "seventh", gift: "seven Swans-a-Swimming" },
    { english: "eighth", gift: "eight Maids-a-Milking" },
    { english: "ninth", gift: "nine Ladies Dancing" },
    { english: "tenth", gift: "ten Lords-a-Leaping" },
    { english: "eleventh", gift: "eleven Pipers Piping" },
    { english: "twelfth", gift: "twelve Drummers Drumming" },
  ]

  def self.song
    12.times.map { |day|
      gifts = day.downto(0).map { |d| @days[d][:gift] }
      if day > 0
        gifts[-1] = "and " + gifts[-1]
      end
      "On the #{@days[day][:english]} day of Christmas my true love gave to me: #{gifts.join(", ")}.\n"
    }.join("\n")
  end
end



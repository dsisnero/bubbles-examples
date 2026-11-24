require "./spec_helper"
ENV["BUBBLETEA_EXAMPLE_DISABLE_MAIN"] = "1"
require "../src/textinputs"

# Init is nil so the example's textinput.Blink tick never fires: a blink tick
# changes the virtual cursor's cell style, and its phase against the renderer's
# frame ticks makes the capture depend on scheduling (same wrapper as
# cellbuffer, splash, stopwatch and timer).
private class TextinputsParityModel < TextinputsModel
  def init : Tea::Cmd?
    nil
  end
end

private def capture_textinputs_output : Bytes
  output = IO::Memory.new
  program = Bubbletea.new_program(
    TextinputsParityModel.new,
    Tea.with_input(IO::Memory.new("")),
    Tea.with_output(output),
    Tea.without_signals,
    Tea.with_window_size(80, 24),
  )

  spawn do
    # Each phase is separated by more than a frame so a burst of keys cannot
    # coalesce into a variable number of frames.
    sleep 100.milliseconds
    program.send(Tea.key('h'))
    program.send(Tea.key('i'))
    sleep 80.milliseconds
    program.send(Tea.key(Tea::KeyTab))
    sleep 80.milliseconds

    program.send(Tea.key('a'))
    program.send(Tea.key('@'))
    program.send(Tea.key('b'))
    sleep 80.milliseconds
    program.send(Tea.key(Tea::KeyTab))
    sleep 80.milliseconds

    program.send(Tea.key('1'))
    program.send(Tea.key('2'))
    sleep 80.milliseconds
    program.send(Tea.key('r', Tea::ModCtrl))
    sleep 80.milliseconds

    program.send(Tea.key(Tea::KeyTab))
    sleep 80.milliseconds
    program.send(Tea.key(Tea::KeyEnter))
  end

  _model, err = program.run
  raise err.not_nil! if err
  output.to_slice
end

describe "examples/textinputs parity" do
  it "matches the saved Go golden output exactly" do
    actual = capture_textinputs_output
    expected = File.read("#{__DIR__}/golden/textinputs.go.golden").to_slice
    actual.should eq(expected)
  end
end

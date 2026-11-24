require "./spec_helper"
require "bubbles"
ENV["BUBBLETEA_EXAMPLE_DISABLE_MAIN"] = "1"
require "../src/dynamic_textarea"

private def capture_dynamic_textarea_output : Bytes
  output = IO::Memory.new
  program = Bubbletea.new_program(
    DynamicTextareaModel.new,
    Tea.with_input(IO::Memory.new("")),
    Tea.with_output(output),
    Tea.without_signals,
    Tea.with_window_size(80, 24),
  )

  spawn do
    sleep 100.milliseconds
    program.send(Tea.key('a'))
    program.send(Tea.key(Tea::KeyEnter))
    program.send(Tea.key('b'))
    program.send(Tea.key(Tea::KeyEnter))
    program.send(Tea.key('c'))
    # The example only quits on ctrl+c, and typing into the textarea would
    # change its content, so quit the way the renderer specs do.
    sleep 120.milliseconds
    program.send(Tea::QuitMsg.new)
  end

  _model, err = program.run
  raise err.not_nil! if err
  output.to_slice
end

describe "examples/dynamic_textarea parity" do
  it "matches the saved Go golden output exactly" do
    actual = capture_dynamic_textarea_output
    expected = File.read("#{__DIR__}/golden/dynamic_textarea.go.golden").to_slice
    actual.should eq(expected)
  end
end

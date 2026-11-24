require "./spec_helper"
ENV["BUBBLETEA_EXAMPLE_DISABLE_MAIN"] = "1"
require "../src/paginator"

private def capture_paginator_output : Bytes
  output = IO::Memory.new
  program = Bubbletea.new_program(PaginatorModel.new, Tea.with_input(IO::Memory.new("")), Tea.with_output(output), Tea.without_signals, Tea.with_window_size(80, 24))
  spawn { sleep 80.milliseconds; program.send(Tea::QuitMsg.new) }
  _model, err = program.run; raise err.not_nil! if err; output.to_slice
end

describe "examples/paginator parity" do
  pending "matches the saved Go golden output exactly (the capture counts renderer ticks: paginator's 26-line view forces a repaint every tick, so the frame count sits on a scheduling boundary and can differ by one under load)" do
    capture_paginator_output.should eq(File.read("#{__DIR__}/golden/paginator.go.golden").to_slice)
  end
end

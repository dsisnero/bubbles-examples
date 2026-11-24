require "./spec_helper"
ENV["BUBBLETEA_EXAMPLE_DISABLE_MAIN"] = "1"
require "../src/program_status"

private def capture_program_status_output : Bytes
  output = IO::Memory.new
  program = Bubbletea.new_program(
    ProgramStatusExampleModel.new,
    Tea.with_input(IO::Memory.new("")),
    Tea.with_output(output),
    Tea.without_signals,
    Tea.with_window_size(80, 24),
  )

  spawn do
    sleep 120.milliseconds
    program.send(Tea.key(Tea::KeyRight))
    sleep 440.milliseconds
    program.send(Tea.key(Tea::KeyLeft))
    sleep 200.milliseconds
    program.send(Tea.key('q'))
  end

  _model, err = program.run
  raise err.not_nil! if err
  output.to_slice
end

describe "examples/program_status parity" do
  it "matches the saved Go golden output exactly" do
    actual = capture_program_status_output
    expected = File.read("#{__DIR__}/golden/program_status.go.golden").to_slice
    actual.should eq(expected)
  end

  it "reports each state with the expected OSC 7501 sequence" do
    output = String.new(capture_program_status_output)

    output.should contain(Ansi.set_program_status(Tea::ProgramStatus.new(
      state: Tea::ProgramStateWorking,
      app: "program-status-example",
      kind: Tea::ProgramStatusKindPermission,
      progress: 10,
      has_progress: true,
      message: "Building the project",
    )))
    output.should contain(Ansi.set_program_status(Tea::ProgramStatus.new(
      state: Tea::ProgramStateIdle,
      app: "program-status-example",
      kind: Tea::ProgramStatusKindPermission,
      message: "Waiting for instructions",
    )))
  end
end

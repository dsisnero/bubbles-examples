require "bubbletea"

struct ProgramStatusTickMsg
  include Tea::Msg
end

# Port of upstream examples/program-status: cycles through the Program Status
# Protocol states and reports them with OSC 7501.
class ProgramStatusExampleModel
  include Bubbletea::Model

  STATES = [
    Tea::ProgramStateIdle,
    Tea::ProgramStateWorking,
    Tea::ProgramStateBlocked,
    Tea::ProgramStateDone,
    Tea::ProgramStateError,
  ]

  MESSAGES = {
    Tea::ProgramStateIdle    => "Waiting for instructions",
    Tea::ProgramStateWorking => "Building the project",
    Tea::ProgramStateBlocked => "Deploy to production?",
    Tea::ProgramStateDone    => "Build finished",
    Tea::ProgramStateError   => "Build failed",
  }

  def initialize
    @state = 0
    @progress = 0
    @supported = false
  end

  def init : Bubbletea::Cmd?
    Tea.batch(Tea.request_program_status_support, tick)
  end

  def update(msg : Tea::Msg)
    case msg
    when Tea::ProgramStatusSupportMsg
      @supported = true
    when ProgramStatusTickMsg
      @progress = (@progress + 5) % 105 if STATES[@state] == Tea::ProgramStateWorking
      return {self, tick}
    when Tea::KeyPressMsg
      case msg.keystroke
      when "q", "ctrl+c"
        return {self, Tea.quit}
      when "right", "l", "space"
        @state = (@state + 1) % STATES.size
        @progress = 0
      when "left", "h"
        @state = (@state + STATES.size - 1) % STATES.size
        @progress = 0
      end
    end

    {self, nil}
  end

  def view : Bubbletea::View
    state = STATES[@state]
    support = @supported ? "supported" : "no reply yet (unsupported, or behind tmux)"

    v = Tea.new_view(
      "Program Status Protocol (OSC 7501): #{support}\n\n" \
      "State: #{state}\n" \
      "Message: #{MESSAGES[state]}\n\n" \
      "Press left/right to change state, q to quit.\n"
    )
    v.program_status = Tea::ProgramStatus.new(
      state: state,
      app: "program-status-example",
      kind: Tea::ProgramStatusKindPermission,
      progress: @progress,
      has_progress: state == Tea::ProgramStateWorking,
      message: MESSAGES[state],
    )
    v
  end

  private def tick : Bubbletea::Cmd
    Tea.tick(200.milliseconds) { ProgramStatusTickMsg.new }
  end
end

unless ENV["BUBBLETEA_EXAMPLE_DISABLE_MAIN"]? == "1"
  program = Bubbletea::Program.new(ProgramStatusExampleModel.new)
  _model, err = program.run
  STDERR.puts "Error: #{err.message}" if err
end

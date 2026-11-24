require "bubbles"
require "bubbletea"

# Port of upstream examples/textinputs: three Bubbles text-input components
# with focus cycling, an echo-mode password field, a cursor-mode toggle and a
# submit button.
class TextinputsModel
  include Bubbletea::Model

  FOCUSED_STYLE          = Lipgloss::Style.new.foreground(Lipgloss.color("205"))
  BLURRED_STYLE          = Lipgloss::Style.new.foreground(Lipgloss.color("240"))
  CURSOR_STYLE           = FOCUSED_STYLE
  NO_STYLE               = Lipgloss::Style.new
  HELP_STYLE             = BLURRED_STYLE
  CURSOR_MODE_HELP_STYLE = Lipgloss::Style.new.foreground(Lipgloss.color("244"))

  FOCUSED_BUTTON = FOCUSED_STYLE.render("[ Submit ]")
  BLURRED_BUTTON = "[ #{BLURRED_STYLE.render("Submit")} ]"

  getter focus_index : Int32
  getter inputs : Array(Bubbles::TextInput::Model)
  getter cursor_mode : Bubbles::Cursor::Mode
  getter? quitting : Bool

  def initialize
    @focus_index = 0
    @inputs = Array(Bubbles::TextInput::Model).new(3)
    @cursor_mode = Bubbles::Cursor::CursorBlink
    @quitting = false

    3.times do |i|
      t = Bubbles::TextInput.new
      t.char_limit = 32

      s = t.styles
      s.cursor.color = Lipgloss.color("205")
      s.focused.prompt = FOCUSED_STYLE
      s.focused.text = FOCUSED_STYLE
      s.blurred.prompt = BLURRED_STYLE
      # Upstream sets Focused.Text twice and leaves Blurred.Text alone; kept so
      # the port styles the same cells.
      s.focused.text = FOCUSED_STYLE
      t.styles = s

      case i
      when 0
        t.placeholder = "Nickname"
        t.focus
      when 1
        t.placeholder = "Email"
        t.char_limit = 64
      when 2
        t.placeholder = "Password"
        t.echo_mode = Bubbles::TextInput::EchoPassword
        t.echo_character = '•'
      end

      @inputs << t
    end
  end

  def init : Bubbletea::Cmd?
    -> { Bubbles::TextInput.blink.as(Tea::Msg?) }
  end

  def update(msg : Tea::Msg)
    case msg
    when Bubbletea::KeyPressMsg
      case msg.keystroke
      when "ctrl+c", "esc"
        @quitting = true
        return {self, Bubbletea.quit}
      when "ctrl+r"
        # Change cursor mode.
        @cursor_mode = cycle_cursor_mode(@cursor_mode)
        cmds = Array(Tea::Cmd?).new(@inputs.size)
        @inputs.each do |input|
          s = input.styles
          s.cursor.blink = @cursor_mode == Bubbles::Cursor::CursorBlink
          input.styles = s
        end
        return {self, Tea.batch(cmds)}
      when "tab", "shift+tab", "enter", "up", "down"
        key = msg.keystroke

        # Did the user press enter while the submit button was focused?
        if key == "enter" && @focus_index == @inputs.size
          return {self, Bubbletea.quit}
        end

        if key == "up" || key == "shift+tab"
          @focus_index -= 1
        else
          @focus_index += 1
        end

        if @focus_index > @inputs.size
          @focus_index = 0
        elsif @focus_index < 0
          @focus_index = @inputs.size
        end

        cmds = Array(Tea::Cmd?).new(@inputs.size)
        @inputs.each_with_index do |input, i|
          if i == @focus_index
            cmds << input.focus
            next
          end

          input.blur
        end

        return {self, Tea.batch(cmds)}
      end
    end

    {self, update_inputs(msg)}
  end

  # update_inputs fans the message out to every input; only focused inputs
  # respond.
  private def update_inputs(msg : Tea::Msg) : Tea::Cmd?
    cmds = Array(Tea::Cmd?).new(@inputs.size)
    @inputs.size.times do |i|
      @inputs[i], cmd = @inputs[i].update(msg)
      cmds << cmd
    end
    Tea.batch(cmds)
  end

  # Upstream cycles iota order Blink -> Static -> Hide -> Blink.
  private def cycle_cursor_mode(mode : Bubbles::Cursor::Mode) : Bubbles::Cursor::Mode
    next_mode = Bubbles::Cursor::Mode.new(mode.value + 1)
    next_mode > Bubbles::Cursor::CursorHide ? Bubbles::Cursor::CursorBlink : next_mode
  end

  def view : Bubbletea::View
    b = String::Builder.new
    cursor = nil.as(Tea::Cursor?)

    @inputs.each_with_index do |input, i|
      b << input.view
      b << '\n' if i < @inputs.size - 1
      if @cursor_mode != Bubbles::Cursor::CursorHide && input.focused?
        if c = input.cursor
          c.y += i
          cursor = c
        end
      end
    end

    button = @focus_index == @inputs.size ? FOCUSED_BUTTON : BLURRED_BUTTON
    b << "\n\n" << button << "\n\n"

    b << HELP_STYLE.render("cursor mode is ")
    b << CURSOR_MODE_HELP_STYLE.render(@cursor_mode.to_s)
    b << HELP_STYLE.render(" (ctrl+r to change style)")

    b << '\n' if @quitting

    v = Bubbletea::View.new(b.to_s)
    v.cursor = cursor
    v
  end
end

unless ENV["BUBBLETEA_EXAMPLE_DISABLE_MAIN"]? == "1"
  program = Bubbletea::Program.new(TextinputsModel.new)
  _model, err = program.run
  if err
    STDERR.puts "Error running program: #{err.message}"
    exit 1
  end
end

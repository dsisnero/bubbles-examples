require "bubbles"
require "bubbletea"

# Port of upstream examples/dynamic-textarea: a textarea that grows with its
# content between a minimum and maximum height.
class DynamicTextareaModel
  include Bubbletea::Model

  property textarea : Bubbles::Textarea::Model

  def initialize
    ti = Bubbles::Textarea.new
    ti.placeholder = "Schnrr..."
    ti.set_show_line_numbers(true)
    ti.dynamic_height = true
    ti.min_height = 3
    ti.max_height = 15
    ti.max_content_height = 20
    ti.set_width(60)
    ti.set_virtual_cursor(false)
    ti.focus
    @textarea = ti
  end

  def init : Bubbletea::Cmd?
    Tea.batch(Bubbles::Textarea::Model.blink, Tea.request_background_color)
  end

  def update(msg : Tea::Msg)
    cmds = [] of Tea::Cmd?
    cmd = nil.as(Tea::Cmd?)

    case msg
    when Tea::BackgroundColorMsg
      @textarea.set_styles(Bubbles::Textarea.default_styles(msg.is_dark?))
    when Tea::KeyPressMsg
      return {self, Tea.quit} if msg.keystroke == "ctrl+c"
    end

    @textarea, cmd = @textarea.update(msg)
    cmds << cmd
    {self, Tea.batch(cmds)}
  end

  # The status line upstream builds with fmt.Sprintf("%.0f", ...*100).
  def status_view : String
    scroll = (@textarea.viewport.scroll_percent * 100).round.to_i
    "\nHeight: #{@textarea.height} · Lines: #{@textarea.line_count} · " \
    "Cursor: (#{@textarea.line}, #{@textarea.column}) · Scroll: #{scroll}%"
  end

  def view : Bubbletea::View
    gap = 1

    cursor = nil.as(Tea::Cursor?)
    unless @textarea.virtual_cursor?
      if c = @textarea.cursor
        c.y += gap
        cursor = c
      end
    end

    body = "\n" + [@textarea.view, status_view, "\n(ctrl+c to quit)"].join("\n")
    v = Bubbletea::View.new(body)
    v.cursor = cursor
    v
  end
end

unless ENV["BUBBLETEA_EXAMPLE_DISABLE_MAIN"]? == "1"
  program = Bubbletea::Program.new(DynamicTextareaModel.new)
  _model, err = program.run
  STDERR.puts "Error running program: #{err.message}" if err
end

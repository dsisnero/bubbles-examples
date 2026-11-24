require "bubbles"

class ViewportExampleModel
  include Tea::Model

  @viewport : Bubbles::Viewport::Model
  @content : String

  def initialize
    @content = (1..60).map { |i| "Line %02d — scroll with j/k, ↑/↓, or page keys" % i }.join('\n')
    @viewport = Bubbles::Viewport.new(
      Bubbles::Viewport.with_width(76),
      Bubbles::Viewport.with_height(18)
    )
    @viewport.set_content(@content)
  end

  def init : Tea::Cmd?
    nil
  end

  def update(msg : Tea::Msg)
    case msg
    when Tea::WindowSizeMsg
      @viewport.set_width(msg.width)
      @viewport.set_height(msg.height - 2)
    when Tea::KeyPressMsg
      key = msg.string
      return {self, Tea.quit} if key == "q" || key == "esc" || key == "ctrl+c"
    end

    @viewport, cmd = @viewport.update(msg)
    {self, cmd}
  end

  def view : Bubbletea::View
    Bubbletea::View.new("Viewport demo — q to quit\n" + @viewport.view)
  end
end

unless ENV["BUBBLETEA_EXAMPLE_DISABLE_MAIN"]? == "1"
  program = Bubbletea::Program.new(ViewportExampleModel.new)
  _model, err = program.run
  if err
    STDERR.puts err.message
    exit 1
  end
end

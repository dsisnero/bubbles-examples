require "bubbles"

class TreeExampleModel
  include Tea::Model

  @tree : Bubbles::Tree::Model

  def initialize
    root = Bubbles::Tree::Node.root("~/charm").child(
      "ayman",
      Bubbles::Tree::Node.root("bash").child(
        Bubbles::Tree::Node.root("tools").child("zsh", "doom-emacs")
      ),
      Bubbles::Tree::Node.root("carlos").child(
        Bubbles::Tree::Node.root("emotes").child("chefkiss.png", "kekw.png")
      ),
      "maas"
    )

    @tree = Bubbles::Tree.new(root, 70, 13)
  end

  def init : Tea::Cmd?
    nil
  end

  def update(msg : Tea::Msg)
    case msg
    when Tea::WindowSizeMsg
      @tree.set_size(msg.width, msg.height)
    when Tea::KeyPressMsg
      key = msg.string
      return {self, Tea.quit} if key == "q" || key == "esc" || key == "ctrl+c"
    end

    @tree, cmd = @tree.update(msg)
    {self, cmd}
  end

  def view : Bubbletea::View
    Bubbletea::View.new(@tree.view)
  end
end

unless ENV["BUBBLETEA_EXAMPLE_DISABLE_MAIN"]? == "1"
  program = Bubbletea::Program.new(TreeExampleModel.new)
  _model, err = program.run
  if err
    STDERR.puts err.message
    exit 1
  end
end

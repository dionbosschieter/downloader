require 'find'

class Movie
  def initialize(path)
    @path = path
  end
  def files
    if is_movie_file(@path)
      return [Moviefile.new(@path)]
    end
    if File.file?(@path)
      return []
    end

    files = []
    Find.find(@path) { |file|
      if is_movie_file(file)
        files.push(Moviefile.new(file))
      end
    }

    return files
  end

  def is_movie_file(file)
    if File.file?(file) == false
      return false
    end
    return file =~ /\.avi$|\.mkv$|\.mp4$/
  end
end

class Moviefile
  def initialize(path)
    @path = path
  end

  def filename
    return @path
  end

  def can_be_skipped?
    return File.file?(subtitle_path)
  end

  def subtitle_path
    return @path.gsub(/\.avi$|\.mkv$|\.mp4$/, '.srt')
  end

  def get_subcmd
      return ["/home/dion/go/bin/subify", "dl", @path, "-l", "nl,en"].join(' ')
  end
end

class Application
  def initialize(path)
    @path = path
    @folders = Dir.entries(path)
  end

  def fire
    @folders.each { |folder|
      if folder =~ /^\./
        next
      end
      movie = Movie.new(@path + '/' + folder)

      movie.files.each { |file|
        puts '* scanning for ' + file.filename
        if file.can_be_skipped?
          puts " skipped, srt file already there\n"
        end

        command = file.get_subcmd
        puts command
        system(command)
     }
    }
  end
end

app = Application.new(ARGV.join(''))
app.fire


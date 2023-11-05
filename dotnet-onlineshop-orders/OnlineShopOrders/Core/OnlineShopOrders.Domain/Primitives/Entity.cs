public abstract class Entity:IEquatable<Entity>
{
  public Entity(ulong id)
  {
    Id=id;
  }
  protected Entity()
  {

  }

  public ulong Id{get;protected init;}
  public bool Equals(Entity? other)
  {
    if(other is null)
    return false;

    if(other.GetType() !=this.GetType())
    return false;

    if(other is not  Entity entity)
    return false;

    return other.Id==Id;
  }
  
  public static bool operator ==(Entity? left,Entity? right)
  {
    return left is not null && right is not null && left.Equals(right);
  }

  public static bool operator !=(Entity? left,Entity? right )
  {
    return !(left==right);
  }

   public override bool Equals(object obj)
    {
        if (obj is null)
            return false;
        if (obj.GetType() != this.GetType())
            return false;
        
        return Equals(obj as Entity);
    }

    public override int GetHashCode()
    {
        return this.GetHashCode() * 41;
    }
}